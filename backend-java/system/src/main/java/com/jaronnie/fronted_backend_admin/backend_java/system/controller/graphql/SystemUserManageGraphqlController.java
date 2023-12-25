package com.jaronnie.fronted_backend_admin.backend_java.system.controller.graphql;

import com.jaronnie.fronted_backend_admin.backend_java.system.domain.query.PageQuery;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.query.SearchUserQuery;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.TableDataInfo;
import com.jaronnie.fronted_backend_admin.backend_java.system.domain.vo.UserVo;
import com.jaronnie.fronted_backend_admin.backend_java.system.service.IUserService;
import lombok.RequiredArgsConstructor;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.web.bind.annotation.*;

@RestController
@RequiredArgsConstructor
public class SystemUserManageGraphqlController {
    private final IUserService iUserService;

    @QueryMapping
    public TableDataInfo<UserVo> getUsers() {
        return iUserService.queryPageList(new PageQuery(), new SearchUserQuery());
    }
}

