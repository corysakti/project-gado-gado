package com.example.demo.service;

import com.example.demo.entity.Product;
import com.example.demo.request.ProductRequest;
import com.example.demo.response.DataResponse;

public interface ProductService {
    DataResponse<?> addProduct(ProductRequest product);
    DataResponse<?> updateProduct(ProductRequest product);
    DataResponse<?> deleteProduct(long productId);
    DataResponse<?> findProductById(long productId);
    DataResponse<?> findAllProduct();
}
