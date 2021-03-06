export interface User {
    name?: string;
    size?: number; // number of players required
    botsInvites?: number,
    tokenId: string;
    room?: string;
    die?: number[];
    shots?: 0 | 1 | 2;
    finish?: boolean;
}
