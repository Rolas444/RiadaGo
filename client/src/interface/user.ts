export interface createUser {
    Username: string;
    password: string;
    role: string;
}

export interface userLoged {
    userId: number;
    username: string;
    rol: string;
    login: boolean;
}