import { withTokenRequest, getRequestHeaders } from "../http";

export const postTeams = async (team: any): Promise<void> => {
    try {
        await withTokenRequest.post('/sv/teams', team, {
            headers: getRequestHeaders()
        });
    } catch (error) {
        console.error('Error posting team:', error);
        throw error;
    }
}