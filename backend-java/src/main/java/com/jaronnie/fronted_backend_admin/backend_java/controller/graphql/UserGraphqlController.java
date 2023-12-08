package com.jaronnie.fronted_backend_admin.backend_java.controller.graphql;

import com.jaronnie.fronted_backend_admin.backend_java.domain.bo.PageQuery;
import com.jaronnie.fronted_backend_admin.backend_java.domain.vo.TableDataInfo;
import com.jaronnie.fronted_backend_admin.backend_java.domain.vo.UserVo;
import com.jaronnie.fronted_backend_admin.backend_java.service.IUserService;
import lombok.RequiredArgsConstructor;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.web.bind.annotation.*;

@RestController
@RequiredArgsConstructor
public class UserGraphqlController {
    private final IUserService iUserService;

    @QueryMapping
    public TableDataInfo<UserVo> getUsers() {
        return iUserService.queryPageList(new PageQuery());
    }
}

