package com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo;

import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.OssFileVo;
import lombok.Data;

@Data
public class AddUserBo {
    String username;
    String password;
    OssFileVo avatar;
}
