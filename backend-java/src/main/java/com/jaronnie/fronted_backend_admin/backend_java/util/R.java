package com.jaronnie.fronted_backend_admin.backend_java.util;

import lombok.Getter;

import java.io.Serializable;

@Getter
public class R<T> implements Serializable {
    private static final long serialVersionUID = 1L;
    public static final int SUCCESS = 200;
    public static final int FAIL = 500;
    private int code;
    private String msg;
    private T data;

    public static <T> R<Object> ok() {
        return restResult(null, 200, null);
    }

    public static <T> R<T> ok(T data) {
        return restResult(data, 200, null);
    }

    public static <T> R<Object> ok(String msg) {
        return restResult(null, 200, msg);
    }

    public static <T> R<T> ok(String msg, T data) {
        return restResult(data, 200, msg);
    }

    public static <T> R<Object> fail() {
        return restResult(null, 500, null);
    }

    public static <T> R<Object> fail(String msg) {
        return restResult(null, 500, msg);
    }

    public static <T> R<T> fail(T data) {
        return restResult(data, 500, null);
    }

    public static <T> R<T> fail(String msg, T data) {
        return restResult(data, 500, msg);
    }

    public static <T> R<Object> fail(int code, String msg) {
        return restResult(null, code, msg);
    }

    private static <T> R<T> restResult(T data, int code, String msg) {
        R<T> r = new R<>();
        r.setCode(code);
        r.setData(data);
        r.setMsg(msg);
        return r;
    }

    public static <T> R<Object> count(Integer count) {
        return count > 0 ? ok() : fail();
    }

    public static <T> R<Object> result(boolean result) {
        return result ? ok() : fail();
    }

    private Boolean isOk() {
        return 200 == this.code;
    }

    public void setCode(int code) {
        this.code = code;
    }

    public void setMsg(String msg) {
        this.msg = msg;
    }

    public void setData(T data) {
        this.data = data;
    }

    public boolean equals(Object o) {
        if (o == this) {
            return true;
        } else if (!(o instanceof R)) {
            return false;
        } else {
            R<?> other = (R)o;
            if (!other.canEqual(this)) {
                return false;
            } else if (this.getCode() != other.getCode()) {
                return false;
            } else {
                Object this$msg = this.getMsg();
                Object other$msg = other.getMsg();
                if (this$msg == null) {
                    if (other$msg != null) {
                        return false;
                    }
                } else if (!this$msg.equals(other$msg)) {
                    return false;
                }

                Object this$data = this.getData();
                Object other$data = other.getData();
                if (this$data == null) {
                    if (other$data != null) {
                        return false;
                    }
                } else if (!this$data.equals(other$data)) {
                    return false;
                }

                return true;
            }
        }
    }

    protected boolean canEqual(Object other) {
        return other instanceof R;
    }

    public int hashCode() {
        int result = 1;
        result = result * 59 + this.getCode();
        Object $msg = this.getMsg();
        result = result * 59 + ($msg == null ? 43 : $msg.hashCode());
        Object $data = this.getData();
        result = result * 59 + ($data == null ? 43 : $data.hashCode());
        return result;
    }

    public String toString() {
        return "R(code=" + this.getCode() + ", msg=" + this.getMsg() + ", data=" + this.getData() + ")";
    }

    public R() {
    }
}
