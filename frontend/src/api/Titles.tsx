import { withTokenRequest, getRequestHeaders } from '../http';
import { Title } from '../entity/title';

export const getTitles = async (): Promise<Title[]> => {
    try {
        const response = await withTokenRequest.get('/titles', {
            headers: getRequestHeaders()
        });
        return response.data.data.titles;
    } catch (error) {
        console.log('error');
        throw error;
    }
}