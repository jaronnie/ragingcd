package com.jaronnie.ragingcd.stdb.system.domain.query;

import cn.hutool.core.util.ObjectUtil;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import lombok.Data;

@Data
public class PageQuery {
    private int pageNum;
    private int pageSize;

    public static final int DEFAULT_PAGE_NUM = 1;
    public static final int DEFAULT_PAGE_SIZE = Integer.MAX_VALUE;

    public <T> Page<T> build() {
        Integer pageNum = ObjectUtil.defaultIfNull(this.getPageNum(), 1);
        Integer pageSize = ObjectUtil.defaultIfNull(this.getPageSize(), Integer.MAX_VALUE);
        if (pageNum <= 0) {
            pageNum = 1;
        }
        if (pageSize <= 0) {
            pageSize = Integer.MAX_VALUE;
        }


        return new Page<>((long) pageNum, (long) pageSize);
    }
}
