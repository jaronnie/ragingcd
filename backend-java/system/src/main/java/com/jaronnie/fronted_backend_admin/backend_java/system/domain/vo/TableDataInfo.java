package com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo;

import com.baomidou.mybatisplus.core.metadata.IPage;
import lombok.Getter;

import java.io.Serializable;
import java.util.List;

@Getter
public class TableDataInfo<T> implements Serializable {
    private static final long serialVersionUID = 1L;
    private long total;
    private List<T> rows;

    public TableDataInfo(List<T> list, long total) {
        this.rows = list;
        this.total = total;
    }

    public static <T> TableDataInfo<T> build(IPage<T> page) {
        TableDataInfo<T> rspData = new TableDataInfo<>();
        rspData.setRows(page.getRecords());
        rspData.setTotal(page.getTotal());
        return rspData;
    }

    public static <T> TableDataInfo<T> build(List<T> list) {
        TableDataInfo<T> rspData = new TableDataInfo<>();
        rspData.setRows(list);
        rspData.setTotal((long) list.size());
        return rspData;
    }

    public static <T> TableDataInfo<T> build() {
        return new TableDataInfo<>();
    }

    public void setTotal(long total) {
        this.total = total;
    }

    public void setRows(List<T> rows) {
        this.rows = rows;
    }

    public boolean equals(Object o) {
        if (o == this) {
            return true;
        } else if (!(o instanceof TableDataInfo)) {
            return false;
        } else {
            TableDataInfo<?> other = (TableDataInfo) o;
            if (!other.canEqual(this)) {
                return false;
            } else if (this.getTotal() != other.getTotal()) {
                return false;
            } else {
                Object this$rows = this.getRows();
                Object other$rows = other.getRows();
                if (this$rows == null) {
                    if (other$rows != null) {
                        return false;
                    }
                } else if (!this$rows.equals(other$rows)) {
                    return false;
                }

                return true;
            }
        }
    }

    protected boolean canEqual(Object other) {
        return other instanceof TableDataInfo;
    }

    public int hashCode() {
        boolean PRIME = true;
        int result = 1;
        long $total = this.getTotal();
        result = result * 59 + (int) ($total >>> 32 ^ $total);
        Object $rows = this.getRows();
        result = result * 59 + ($rows == null ? 43 : $rows.hashCode());
        return result;
    }

    public String toString() {
        return "TableDataInfo(total=" + this.getTotal() + ", rows=" + this.getRows() + ")";
    }

    public TableDataInfo() {
    }
}
