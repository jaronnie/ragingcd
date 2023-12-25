package com.jaronnie.fronted_backend_admin.backend_java.system.domain.bo;

import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.OssFileVo;
import lombok.Data;

@Data
public class AddUserBo {
    private String username;
    private String password;
    private OssFileVo avatar;
    private String email;
}
