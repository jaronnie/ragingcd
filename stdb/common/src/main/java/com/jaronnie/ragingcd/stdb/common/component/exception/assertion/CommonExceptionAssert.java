package com.jaronnie.ragingcd.stdb.common.component.exception.assertion;

import com.jaronnie.ragingcd.stdb.common.component.exception.BaseException;
import com.jaronnie.ragingcd.stdb.common.component.exception.CommonException;
import com.jaronnie.ragingcd.stdb.common.component.exception.IResponseEnum;

import java.text.MessageFormat;

public interface CommonExceptionAssert extends IResponseEnum, Assert {

    @Override
    default BaseException newException(Object... args) {
        String msg = this.getMessage();
        if (args != null && args.length > 0) {
            msg = MessageFormat.format(this.getMessage(), args);
        }

        return new CommonException(this, args, msg);
    }

    @Override
    default BaseException newException(Throwable t, Object... args) {
        String msg = this.getMessage();
        if (args != null && args.length > 0) {
            msg = MessageFormat.format(this.getMessage(), args);
        }

        return new CommonException(this, args, msg, t);
    }
}
