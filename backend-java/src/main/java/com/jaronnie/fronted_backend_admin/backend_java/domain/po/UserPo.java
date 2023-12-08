package com.jaronnie.fronted_backend_admin.backend_java.domain.po;

import com.baomidou.mybatisplus.annotation.FieldStrategy;
import com.baomidou.mybatisplus.annotation.TableField;
import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.*;

import javax.persistence.Entity;
import javax.persistence.Id;
import java.io.Serializable;

@EqualsAndHashCode(callSuper = false)
@TableName(value = "user", autoResultMap = true)
@Entity
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class UserPo extends BaseEntity implements Serializable {
    private static final long serialVersionUID = 1L;

    @Id
    @TableId
    private Integer id;

    @TableField(updateStrategy = FieldStrategy.NOT_EMPTY)
    private String username;
    @TableField(updateStrategy = FieldStrategy.NOT_EMPTY)
    private String password;
}
