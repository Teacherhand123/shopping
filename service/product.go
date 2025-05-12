package service

import (
	"context"
	"mime/multipart"
	"shopping/dao"
	"shopping/model"
	"shopping/pkg/e"
	"shopping/pkg/utils"
	"shopping/serializer"
	"strconv"
	"sync"
)

type ProductService struct {
	Id            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"prirce" form:"prirce"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}

func (service *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.GetUserById(uId)
	// 以第一张作为封面图
	tmp, _ := files[0].Open()
	path, err := uploadProductToLocalStatic(tmp, uId, service.Name)

	if err != nil {
		code = e.ErrorProductImgUpload
		utils.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	product := model.Product{
		Name:          service.Name,
		CategoryId:    service.CategoryId,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           service.Num,
		BossId:        uId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}

	// 添加商品到数据库
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(&product)

	if err != nil {
		code = e.Error
		utils.LogrusObj.Infoln(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 子线程开始上传商品图片
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		// 在 Goroutine 中使用循环变量（如 index 和 file）时，必须将它们作为参数传入 Goroutine，避免闭包导致的变量共享问题。
		go func(index int, file *multipart.FileHeader) {
			defer wg.Done() // 确保 Goroutine 结束时调用 Done()

			num := strconv.Itoa(index)
			productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
			tmp, _ := file.Open()
			path, err := uploadProductToLocalStatic(tmp, uId, service.Name+num)
			if err != nil {
				code = e.ErrorProductImgUpload
				// 错误处理需要同步到主线程，不能直接 return
				utils.LogrusObj.Errorln(err)
				return
			}

			productImg := model.ProductImg{
				ProductId: product.ID,
				ImgPath:   path,
			}

			err = productImgDao.CreateProductImg(&productImg)
			if err != nil {
				code = e.Error
				utils.LogrusObj.Errorln(err)
				return
			}
		}(index, file) // 将循环变量传入 Goroutine，避免闭包问题
	}
	wg.Wait()

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(&product),
	}
}

func (service *ProductService) List(ctx context.Context) serializer.Response {
	var products []*model.Product
	var err error
	code := e.Success

	// 默认最多拿15条记录
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	condition := make(map[string]interface{})
	if service.CategoryId != 0 {
		condition["category_id"] = service.CategoryId
	}
	productDao := dao.NewProductDao(ctx)
	total, err := productDao.CountProductByCondition(condition)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		// 如果 productDao.DB 绑定了 ctx，多个 Goroutine 会共享同一个 gorm.DB 实例。
		// 如果 ctx 被取消或超时，可能会影响多个 Goroutine 的操作。
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, service.BasePage)
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

func (service *ProductService) Search(ctx context.Context) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	productDao := dao.NewProductDao(ctx)
	products, count, err := productDao.SearchProduct(service.Info, service.BasePage)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(count))
}

func (service *ProductService) Show(ctx context.Context, id string) serializer.Response {
	code := e.Success
	pId, _ := strconv.Atoi(id)
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(uint(pId))

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}
