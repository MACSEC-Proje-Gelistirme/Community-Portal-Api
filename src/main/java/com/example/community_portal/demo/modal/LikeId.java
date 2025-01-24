package com.example.community_portal.demo.modal;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.UUID;

// Composite Key for Likes
@Data
@NoArgsConstructor
@AllArgsConstructor
public class LikeId {
    private String user;
    private UUID post;
}
