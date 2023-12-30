package com.jaronnie.ragingcd.stdb.common.component.handle;

import cn.dev33.satoken.exception.NotLoginException;
import com.fasterxml.jackson.databind.exc.InvalidFormatException;
import com.jaronnie.ragingcd.stdb.common.component.exception.*;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.ConversionNotSupportedException;
import org.springframework.beans.TypeMismatchException;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.http.converter.HttpMessageNotReadableException;
import org.springframework.http.converter.HttpMessageNotWritableException;
import org.springframework.stereotype.Component;
import org.springframework.validation.BindException;
import org.springframework.validation.BindingResult;
import org.springframework.validation.FieldError;
import org.springframework.validation.ObjectError;
import org.springframework.web.HttpMediaTypeNotAcceptableException;
import org.springframework.web.HttpMediaTypeNotSupportedException;
import org.springframework.web.HttpRequestMethodNotSupportedException;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.MissingPathVariableException;
import org.springframework.web.bind.MissingServletRequestParameterException;
import org.springframework.web.bind.ServletRequestBindingException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.context.request.async.AsyncRequestTimeoutException;
import org.springframework.web.method.annotation.MethodArgumentTypeMismatchException;
import org.springframework.web.multipart.support.MissingServletRequestPartException;
import org.springframework.web.servlet.NoHandlerFoundException;

@Slf4j
@Component
@ControllerAdvice
public class GlobalExceptionHandler {
    /**
     * 处理参数异常，对入参进行检查
     *
     * @param e
     * @return
     */
    @ExceptionHandler(MethodArgumentNotValidException.class)
    @ResponseBody
    public Response<?> validExceptionHandler(MethodArgumentNotValidException e) {
        log.error("参数绑定校验异常", e);
        return wrapperBindingResult(e.getBindingResult());
    }

    /**
     * 业务层异常处理
     *
     * @param exception
     * @return
     */
    @ExceptionHandler(value = BusinessException.class)
    @ResponseBody
    public Response<?> businessExceptionHandler(BaseException exception) {
        log.error("business exception! message:", exception);
        return Response.error(exception.getResponseEnum());
    }

    /**
     * 业务层异常处理
     *
     * @param exception e
     * @return Response
     */
    @ExceptionHandler(value = CommonException.class)
    @ResponseBody
    public Response<?> commonExceptionHandler(BaseException exception) {
        log.error("common exception! message:", exception);
        return Response.error(exception.getResponseEnum());
    }

    /**
     * 空指针异常
     *
     * @param exception e
     * @return Response
     */
    @ExceptionHandler(value = NullPointerException.class)
    @ResponseBody
    public Response<?> nullExceptionHandler(NullPointerException exception) {
        log.error("null pointer exception! message:", exception);
        return Response.error(ResponseEnum.NullPointer_ERROR);
    }

    /**
     * 发生未知异常
     *
     * @param e Exception
     * @return Response
     */
    @ExceptionHandler(Exception.class)
    @ResponseBody
    public Response<?> unknownHandleException(Exception e) {
        log.error("unknown exception! message:", e);
        return Response.error(ResponseEnum.INTERNAL_SERVER_ERROR);
    }

    /**
     * Controller 上一层相关的异常
     *
     * @param e exception
     * @return Response
     */
    @ExceptionHandler({
            NoHandlerFoundException.class,
            HttpRequestMethodNotSupportedException.class,
            HttpMediaTypeNotSupportedException.class,
            MissingPathVariableException.class,
            MissingServletRequestParameterException.class,
            TypeMismatchException.class,
            HttpMessageNotReadableException.class,
            HttpMessageNotWritableException.class,
            HttpMediaTypeNotAcceptableException.class,
            ServletRequestBindingException.class,
            ConversionNotSupportedException.class,
            MissingServletRequestPartException.class,
            AsyncRequestTimeoutException.class
    })
    @ResponseBody
    public Response<?> servletHandleException(Exception e) {
        log.error("exception occurred the request did not enter the Controller: ", e);
        int code = ResponseEnum.SERVER_ERROR.getCode();
        String msg = ResponseEnum.SERVER_ERROR.getMessage();
        if (e.getCause() instanceof InvalidFormatException) {
            InvalidFormatException rootCause = (InvalidFormatException) e.getCause();
            String causeMsg = "字段\"" + rootCause.getPath().get(rootCause.getPath().size() - 1).getFieldName() + "\"超出取值范围";
            return Response.error(code, causeMsg);
        }
        try {
            String name = e.getClass().getSimpleName();
            ServletResponseEnum servletExceptionEnum = ServletResponseEnum.valueOf(name);
            code = servletExceptionEnum.getCode();
            msg = servletExceptionEnum.getMessage();
        } catch (IllegalArgumentException e1) {
            log.error("class [{}] not defined in enum {}", e.getClass().getName(), ServletResponseEnum.class.getName());
        }
        return Response.error(code, msg);
    }

    /**
     * 处理绑定异常
     */
    @ExceptionHandler(BindException.class)
    @ResponseBody
    public Response<?> validExceptionHandler(BindException e) {
        log.error("参数绑定校验异常: ", e);
        return wrapperBindingResult(e.getBindingResult());
    }

    /**
     * 包装绑定异常结果
     *
     * @param bindingResult 绑定结果
     * @return 异常结果
     */
    private Response<?> wrapperBindingResult(BindingResult bindingResult) {
        StringBuilder msg = new StringBuilder();
        for (ObjectError error : bindingResult.getAllErrors()) {
            msg.append(", ");
            if (error instanceof FieldError) {
                msg.append(((FieldError) error).getField()).append(": ");
            }
            msg.append(error.getDefaultMessage() == null ? "" : error.getDefaultMessage());
        }
        return Response.error(ResponseEnum.VALID_ERROR.getCode(), msg.substring(2));
    }

    /**
     * 处理MethodArgumentTypeMismatchException异常
     */
    @ExceptionHandler(value = MethodArgumentTypeMismatchException.class)
    public ResponseEntity<Object> argTypeMismatchHandler(MethodArgumentTypeMismatchException e) {
        String err = e.getName() + " should be of type " + e.getRequiredType().getName();
        log.error(err);
        Response<?> res = Response.error(ResponseEnum.Method_Argument_Type_Mismatch_ERROR);
        return new ResponseEntity<>(res, new HttpHeaders(), HttpStatus.BAD_REQUEST);
    }

    @ExceptionHandler(NotLoginException.class)
    @ResponseBody
    public Response<?> HandlerNotLoginException(NotLoginException nle) throws Exception {

        // 打印堆栈，以供调试
        nle.printStackTrace();

        // 判断场景值，定制化异常信息
        String message = "";
        if(nle.getType().equals(NotLoginException.NOT_TOKEN)) {
            message = "未能读取到有效 token";
        }
        else if(nle.getType().equals(NotLoginException.INVALID_TOKEN)) {
            message = "token 无效";
        }
        else if(nle.getType().equals(NotLoginException.TOKEN_TIMEOUT)) {
            message = "token 已过期";
        }
        else if(nle.getType().equals(NotLoginException.BE_REPLACED)) {
            message = "token 已被顶下线";
        }
        else if(nle.getType().equals(NotLoginException.KICK_OUT)) {
            message = "token 已被踢下线";
        }
        else if(nle.getType().equals(NotLoginException.TOKEN_FREEZE)) {
            message = "token 已被冻结";
        }
        else if(nle.getType().equals(NotLoginException.NO_PREFIX)) {
            message = "未按照指定前缀提交 token";
        }
        else {
            message = "当前会话未登录";
        }

        // 返回给前端
        return Response.error(401, message);
    }
}
