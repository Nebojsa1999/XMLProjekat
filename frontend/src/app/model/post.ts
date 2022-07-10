import { User } from "./user";

export interface Post{
    id: string;          
	ownerId: string;       
	content: string;       
	image: string;         
	likesCount: number;    
	dislikesCount: number; 
	comments: Comment[];      
	link: string[];          
	whoLiked: string[];      
	whoDisliked: string[];   
	user: User;
	postedAt: Date;   
	
}

// export interface T{
// 	posts: Post[];
// }