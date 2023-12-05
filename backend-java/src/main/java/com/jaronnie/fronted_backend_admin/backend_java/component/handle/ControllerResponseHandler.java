package com.jaronnie.fronted_backend_admin.backend_java.component.handle;

import com.jaronnie.fronted_backend_admin.backend_java.component.exception.Response;
import lombok.extern.slf4j.Slf4j;
import org.springframework.core.MethodParameter;
import org.springframework.http.MediaType;
import org.springframework.http.converter.HttpMessageConverter;
import org.springframework.http.server.ServerHttpRequest;
import org.springframework.http.server.ServerHttpResponse;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.mvc.method.annotation.ResponseBodyAdvice;

/**
 * 全局响应结果（ResponseBody）处理器
 */
@Slf4j
@Component
public class ControllerResponseHandler implements ResponseBodyAdvice<Object> {
    /**
     * beforeBodyWrite
     *
     * @param data       the body to be written
     * @param returnType the return type of the controller method
     * @param mt         the content type selected through content negotiation
     * @param aClass     the converter type selected to write to the response
     * @param request    the current request
     * @param response   the current response
     * @return
     */
    @Override
    @SuppressWarnings("NullableProblems")
    public Object beforeBodyWrite(Object data, MethodParameter returnType, MediaType mt,
                                  Class<? extends HttpMessageConverter<?>> aClass, ServerHttpRequest request, ServerHttpResponse response) {
        final String returnTypeName = returnType.getParameterType().getName();
        String voidType = "void";
        if (voidType.equals(returnTypeName)) {
            return Response.success(null);
        }
        if (returnTypeName.equals("org.springframework.http.ResponseEntity")) {
            return data;
        }
        return Response.success(data);
    }

    /**
     * 需处理的返回类型
     *
     * @param returnType the return type
     * @param arg1       the selected converter type
     * @return
     */
    @Override
    @SuppressWarnings("NullableProblems")
    public boolean supports(MethodParameter returnType, Class<? extends HttpMessageConverter<?>> arg1) {
        if (returnType.getMethod() == null) {
            return false;
        }
        return !returnType.getGenericParameterType().equals(Response.class);
    }
}
