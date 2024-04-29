package handler

import (
	"context"
	"git.imooc.com/cap1573/product/common"
	"git.imooc.com/cap1573/product/domain/model"
	"git.imooc.com/cap1573/product/domain/service"
	. "git.imooc.com/cap1573/product/proto/product"
)

type Product struct {
	ProductDataService service.IProductDataService
}

// 创建商品
func (h *Product) AddProduct(ctx context.Context, request *ProductInfo, response *ResponseProduct) error {
	// 函数签名说明：
	// - ctx: 上下文对象，用于携带请求的元数据和取消信号等。
	// - request: 请求参数，这里是一个 ProductInfo 对象，包含了要添加的产品信息。
	// - response: 响应参数，这里是一个 ResponseProduct 对象，用于返回产品 ID。
	// - error: 错误类型，用于指示方法是否成功执行以及出现的任何错误。

	// 创建一个空的 Product 对象，用于保存从请求中提取的产品信息。
	productAdd := &model.Product{}

	// 使用 common 包中的 SwapTo 函数，将请求中的数据转换为 Product 对象。
	if err := common.SwapTo(request, productAdd); err != nil {
		return err
	}

	// 调用 ProductDataService 接口的 AddProduct 方法，将产品信息保存到数据库中，并获取产品 ID。
	productID, err := h.ProductDataService.AddProduct(productAdd)
	if err != nil {
		return err
	}

	// 将获取的产品 ID 设置到响应参数中，以便返回给客户端。
	response.ProductId = productID

	// 返回 nil，表示方法执行成功。
	return nil
}

// 根据id查找商品
func (h *Product) FindProductByID(ctx context.Context, request *RequestID, response *ProductInfo) error {
	productData, err := h.ProductDataService.FindProductByID(request.ProductId)
	if err != nil {
		return err
	}
	if err := common.SwapTo(productData, response); err != nil {
		return err
	}
	return nil
}

// 商品更新
func (h *Product) UpdateProduct(ctx context.Context, request *ProductInfo, response *Response) error {
	productAdd := &model.Product{}
	if err := common.SwapTo(request, productAdd); err != nil {
		return err
	}
	err := h.ProductDataService.UpdateProduct(productAdd)
	if err != nil {
		return err
	}
	response.Msg = "更新成功"
	return nil
}

// 商品删除
func (h *Product) DeleteProductByID(ctx context.Context, request *RequestID, response *Response) error {
	if err := h.ProductDataService.DeleteProduct(request.ProductId); err != nil {
		return err
	}
	response.Msg = "删除成功"
	return nil

}

// 查询所有商品
func (h *Product) FindAllProduct(ctx context.Context, request *RequestAll, response *AllProduct) error {
	productAll, err := h.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}

	for _, v := range productAll {
		productInfo := &ProductInfo{}
		err := common.SwapTo(v, productInfo)
		if err != nil {
			return err
		}
		response.ProductInfo = append(response.ProductInfo, productInfo)
	}
	return nil
}
