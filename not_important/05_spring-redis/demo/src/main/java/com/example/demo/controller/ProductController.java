package com.example.demo.controller;

import com.example.demo.entity.Product;
import com.example.demo.request.ProductRequest;
import com.example.demo.response.DataResponse;
import com.example.demo.service.ProductService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cache.annotation.CacheEvict;
import org.springframework.cache.annotation.CachePut;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ProductController {
    @Autowired
    ProductService productService;

    /*
        @Cacheable is employed to fetch data from the database, storing it in the cache. Upon future invocations,
        the method retrieves the cached value directly, eliminating the need to execute the method again.
     */
    @GetMapping("/product/{id}")
    @Cacheable(cacheNames = "product", key = "#id")
    public DataResponse<?> getProductById(@PathVariable long id) {
        return productService.findProductById(id);
    }

    /*
    @CachePut is used to update data in the cache when there is any update in the source database.
     */
    @PutMapping("/product/{id}")
    @CachePut(cacheNames = "product", key = "#id")
    public DataResponse<?> editProduct(@PathVariable long id, @RequestBody ProductRequest request) {
        request.setId(id);
        return productService.updateProduct(request);
    }

    /*
    @CacheEvict is used for removing stale or unused data from the cache.
    The beforeInvocation attribute allows us to control the eviction process,
    enabling us to choose whether the eviction should occur before or after the method execution.
     */
    @DeleteMapping("/product/{id}")
    @CacheEvict(cacheNames = "product", key = "#id", beforeInvocation = true)
    public DataResponse<?> removeProductById(@PathVariable long id) {
        return productService.deleteProduct(id);
    }
}
