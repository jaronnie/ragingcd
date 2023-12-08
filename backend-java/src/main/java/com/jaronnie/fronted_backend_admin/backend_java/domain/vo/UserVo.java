package com.jaronnie.fronted_backend_admin.backend_java.domain.vo;

import lombok.Builder;
import lombok.Data;

import java.util.Date;

@Data
@Builder
public class UserVo {
    Integer id;
    String username;
    Date createTime;
    Date updateTime;
}
