package com.jaronnie.fronted_backend_admin.backend_java.domain.vo;

import lombok.Builder;
import lombok.Data;

import java.util.Date;

@Data
@Builder
public class UserVo {
    private Integer id;
    private String username;
    private Date createTime;
    private Date updateTime;
    private String avatar;
}
