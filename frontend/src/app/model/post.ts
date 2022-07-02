export interface Post{
    id: string;          
	ownerId: string;       
	content: string;       
	image: string;         
	likesCount: number;    
	dislikesCount: number; 
	comments: Comment[];      
	links: string[];          
	whoLiked: string[];      
	whoDisliked: string[];   
	postedAt: Date;      
}