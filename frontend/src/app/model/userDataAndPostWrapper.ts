import { Post } from "./post";
import { User } from "./user"

export interface UserDataAndPostWrapper {
    user: User,
    post: Post
}
