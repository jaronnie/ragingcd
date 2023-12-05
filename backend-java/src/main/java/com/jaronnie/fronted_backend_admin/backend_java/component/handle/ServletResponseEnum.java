package com.jaronnie.fronted_backend_admin.backend_java.component.handle;

import lombok.AllArgsConstructor;
import lombok.Getter;

import javax.servlet.http.HttpServletResponse;

@Getter
@AllArgsConstructor
public enum ServletResponseEnum {
    MethodArgumentNotValidException(400, "BAD REQUEST", HttpServletResponse.SC_BAD_REQUEST),
    MethodArgumentTypeMismatchException(400, "BAD REQUEST", HttpServletResponse.SC_BAD_REQUEST),
    MissingServletRequestPartException(400, "BAD REQUEST", HttpServletResponse.SC_BAD_REQUEST),
    MissingPathVariableException(400, "BAD REQUEST", HttpServletResponse.SC_BAD_REQUEST),
    BindException(400, "BAD REQUEST", HttpServletResponse.SC_BAD_REQUEST),
    MissingServletRequestParameterException(400, "BAD REQUEST", HttpServletResponse.SC_BAD_REQUEST),
    TypeMismatchException(400, "BAD REQUEST", HttpServletResponse.SC_BAD_REQUEST),
    ServletRequestBindingException(400, "BAD REQUEST", HttpServletResponse.SC_BAD_REQUEST),
    HttpMessageNotReadableException(400, "BAD REQUEST", HttpServletResponse.SC_BAD_REQUEST),
    NoHandlerFoundException(404, "NOT FOUND", HttpServletResponse.SC_NOT_FOUND),
    NoSuchRequestHandlingMethodException(404, "NOT FOUND", HttpServletResponse.SC_NOT_FOUND),
    HttpRequestMethodNotSupportedException(405, "METHOD NOT ALLOWED", HttpServletResponse.SC_METHOD_NOT_ALLOWED),
    HttpMediaTypeNotAcceptableException(406, "MEDIA TYPE NOT ACCEPTABLE", HttpServletResponse.SC_NOT_ACCEPTABLE),
    HttpMediaTypeNotSupportedException(415, "MEDIA TYPE NOT SUPPORTED", HttpServletResponse.SC_UNSUPPORTED_MEDIA_TYPE),
    ConversionNotSupportedException(500, "INTERNAL SERVER ERROR", HttpServletResponse.SC_INTERNAL_SERVER_ERROR),
    HttpMessageNotWritableException(500, "INTERNAL SERVER ERROR", HttpServletResponse.SC_INTERNAL_SERVER_ERROR),
    AsyncRequestTimeoutException(503, "SERVICE UNAVAILABLE", HttpServletResponse.SC_SERVICE_UNAVAILABLE);

    /**
     * 返回码，目前与{@link #statusCode}相同
     */
    private final int code;
    /**
     * 返回信息，直接读取异常的message
     */
    private final String message;
    /**
     * HTTP状态码
     */
    private final int statusCode;
}
