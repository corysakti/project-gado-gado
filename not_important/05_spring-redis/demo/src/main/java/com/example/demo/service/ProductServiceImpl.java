package com.example.demo.service;

import com.example.demo.entity.Product;
import com.example.demo.exception.BusinessException;
import com.example.demo.repository.ProductRepository;
import com.example.demo.request.ProductRequest;
import com.example.demo.response.DataResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class ProductServiceImpl implements ProductService {
    @Autowired
    private ProductRepository productRepository;

    @Override
    public DataResponse<?> addProduct(ProductRequest product) {
        productRepository.save(new Product(product));
        return new DataResponse<>();
    }

    @Override
    public DataResponse<?> updateProduct(ProductRequest product) {
        Product existProduct = findById(product.getId());
        existProduct.setName(product.getName());
        existProduct.setPrice(product.getPrice());
        existProduct.setQuantity(product.getQuantity());
        existProduct.setCode(product.getCode());
        productRepository.save(existProduct);
        return new DataResponse<>("Product updated successfully", existProduct);
    }

    @Override
    public DataResponse<?> deleteProduct(long productId) {
        Product existProduct = findById(productId);
        productRepository.delete(existProduct);
        return new DataResponse<>("Product deleted successfully");
    }

    @Override
    public DataResponse<?> findProductById(long productId) {
        return new DataResponse<>(findById(productId));
    }

    @Override
    public DataResponse<?> findAllProduct() {
        return new DataResponse<>(productRepository.findAll());
    }

    public Product findById(long productId) {
        var existProduct = productRepository.findById(productId);
        if (existProduct.isEmpty()) {
            throw new BusinessException("Product not found");
        }
        return existProduct.get();
    }
}
