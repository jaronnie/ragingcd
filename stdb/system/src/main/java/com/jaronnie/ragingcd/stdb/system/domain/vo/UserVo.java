package com.jaronnie.ragingcd.stdb.system.domain.vo;

import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.Builder;
import lombok.Data;

import java.util.Date;

@ApiModel("用户 UserVo")
@Data
@Builder
public class UserVo {
    @ApiModelProperty("用户Id")
    private Integer id;
    private String username;
    private Date createTime;
    private Date updateTime;
    private String avatar;
    private String email;
}
