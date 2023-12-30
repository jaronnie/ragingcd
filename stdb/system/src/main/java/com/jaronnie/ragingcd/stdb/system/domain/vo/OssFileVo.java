package com.jaronnie.ragingcd.stdb.system.domain.vo;

import lombok.Builder;
import lombok.Data;

import java.util.Date;

@Data
@Builder
public class OssFileVo {
    private String filename;
    private Date uploadTime;
    private String url;
}
