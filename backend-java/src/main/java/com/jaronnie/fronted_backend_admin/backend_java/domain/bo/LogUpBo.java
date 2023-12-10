package com.jaronnie.fronted_backend_admin.backend_java.domain.bo;

import com.jaronnie.fronted_backend_admin.backend_java.domain.vo.OssFileVo;
import lombok.Data;

@Data
public class LogUpBo {
    String username;
    String password;
    OssFileVo avatar;
}
