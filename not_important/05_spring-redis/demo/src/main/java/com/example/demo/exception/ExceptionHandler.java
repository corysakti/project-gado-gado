package com.example.demo.exception;

import com.example.demo.response.DataResponse;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RestControllerAdvice
public class ExceptionHandler {

    @Autowired
    private HttpServletRequest request;

    @org.springframework.web.bind.annotation.ExceptionHandler(BusinessException.class)
    @ResponseStatus(HttpStatus.INTERNAL_SERVER_ERROR)
    public DataResponse<?> handleBusinessException(BusinessException e){
        System.out.println("BussinessException : "+e.getMessage()+", \n Request Url : "+request.getRequestURL());
        return new DataResponse<>(e.getMessage());
    }
}
