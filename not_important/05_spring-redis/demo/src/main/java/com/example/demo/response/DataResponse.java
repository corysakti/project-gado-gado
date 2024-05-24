package com.example.demo.response;

import java.util.Collections;
import java.util.List;

public class DataResponse<T> {
    private int code;
    private String msg;
    private T data;

    public DataResponse(T data) {
        this.data = data;
        this.code = 200;
        this.msg = "Success";
    }

    public DataResponse(int code, String msg) {
        this.code = code;
        this.msg = msg;
        this.data = (T) Collections.EMPTY_LIST;
    }

    public DataResponse() {
        this.data = (T) Collections.EMPTY_LIST;
        this.code = 200;
        this.msg = "Success";
    }

    public DataResponse(String msg, T data) {
        this.code = 200;
        this.msg = msg;
        this.data = data;
    }

    public DataResponse(String msg) {
        this.msg = msg;
        this.data = (T) Collections.EMPTY_LIST;
        this.code = 200;
    }
}
