package com.jaronnie.ragingcd.stdb.system.domain.bo;

import com.jaronnie.ragingcd.stdb.system.domain.vo.OssFileVo;
import lombok.Data;

@Data
public class AddUserBo {
    private String username;
    private String password;
    private OssFileVo avatar;
    private String email;
}
