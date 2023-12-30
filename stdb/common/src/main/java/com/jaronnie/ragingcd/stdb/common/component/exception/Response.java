package com.jaronnie.ragingcd.stdb.common.component.exception;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.google.gson.Gson;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;

@Data
@NoArgsConstructor
@JsonInclude(JsonInclude.Include.NON_NULL)
public class Response<T> implements Serializable {
    /**
     * 错误码
     */
    private Integer code;
    /**
     * 返回数据
     */
    private T data;
    /**
     * 错误提示，用户可读
     */
    private String message;

    /**
     * 将传入的 result 对象，转换成另外一个泛型结果的对象
     *
     * @param result 传入的 result 对象
     * @param <T>    返回的泛型
     * @return 新的 Response 对象
     */
    public static <T> Response<T> error(Response<?> result) {
        return error(result.getCode(), result.getMessage());
    }

    public Response(T data) {
        this(ResponseEnum.SUCCESS, data);
    }

    public Response(IResponseEnum responseEnum, T data) {
        this.code = responseEnum.getCode();
        this.message = responseEnum.getMessage();
        this.data = data;
    }

    /**
     * 自定义 成功返回信息
     *
     * @param data res
     * @return Response
     */
    public static <T> Response<T> success(T data) {
        Response<T> result = new Response<>();
        result.code = ResponseEnum.SUCCESS.getCode();
        result.data = data;
        result.message = ResponseEnum.SUCCESS.getMessage();
        return result;
    }

    /**
     * 不带业务数据的错误返回
     *
     * @param responseEnum exception enum
     * @return Response
     */
    public static <T> Response<T> error(IResponseEnum responseEnum) {
        int code = ResponseEnum.INTERNAL_SERVER_ERROR.getCode();
        String msg = ResponseEnum.INTERNAL_SERVER_ERROR.getMessage();
        if (responseEnum != null) {
            code = responseEnum.getCode();
        }
        if (responseEnum != null) {
            msg = responseEnum.getMessage();
        }
        return error(code, msg);
    }

    /**
     * 返回用户指定的错误号与错误描述
     *
     * @param code    code
     * @param message msg
     * @return Response
     */
    public static <T> Response<T> error(Integer code, String message) {
        Response<T> result = new Response<>();
        result.code = code;
        result.message = message;
        result.data = null;
        return result;
    }

    /**
     * 带业务数据的错误返回
     *
     * @param responseEnum exception enum
     * @return Response
     */
    public static <T> Response<T> error(IResponseEnum responseEnum, T data) {
        Response<T> response = new Response<>();
        response.setCode(responseEnum.getCode());
        response.setMessage(responseEnum.getMessage());
        response.setData(data);
        return response;
    }

    @Override
    public String toString() {
        return new Gson().toJson(this);
    }
}
