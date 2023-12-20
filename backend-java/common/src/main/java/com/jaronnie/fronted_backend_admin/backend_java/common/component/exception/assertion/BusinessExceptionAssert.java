package com.jaronnie.fronted_backend_admin.backend_java.common.component.exception.assertion;

import com.jaronnie.fronted_backend_admin.backend_java.common.component.exception.BaseException;
import com.jaronnie.fronted_backend_admin.backend_java.common.component.exception.BusinessException;
import com.jaronnie.fronted_backend_admin.backend_java.common.component.exception.IResponseEnum;

import java.text.MessageFormat;

public interface BusinessExceptionAssert extends IResponseEnum, Assert {
    @Override
    default BaseException newException(Object... args) {
        String msg = this.getMessage();
        if (args != null && args.length > 0) {
            msg = MessageFormat.format(this.getMessage(), args);
        }
        return new BusinessException(this, args, msg);
    }

    @Override
    default BaseException newException(Throwable t, Object... args) {
        String msg = this.getMessage();
        if (args != null && args.length > 0) {
            msg = MessageFormat.format(this.getMessage(), args);
        }
        return new BusinessException(this, args, msg, t);
    }
}
