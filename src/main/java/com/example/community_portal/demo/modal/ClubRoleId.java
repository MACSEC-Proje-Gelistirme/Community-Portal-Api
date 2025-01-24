package com.example.community_portal.demo.modal;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.UUID;

// Composite Key for ClubRoles
@Data
@NoArgsConstructor
@AllArgsConstructor
public class ClubRoleId {
    private String user;
    private UUID club;
}
