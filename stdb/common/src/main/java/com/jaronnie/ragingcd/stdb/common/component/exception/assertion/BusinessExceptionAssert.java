package com.jaronnie.ragingcd.stdb.common.component.exception.assertion;

import com.jaronnie.ragingcd.stdb.common.component.exception.BaseException;
import com.jaronnie.ragingcd.stdb.common.component.exception.BusinessException;
import com.jaronnie.ragingcd.stdb.common.component.exception.IResponseEnum;

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
