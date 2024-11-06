import { withTokenRequest, getRequestHeaders } from '../http';
import { Title } from '../entity/Title';

export const getTitles = async (): Promise<Title[]> => {
    try {
        const response = await withTokenRequest.get('/titles', {
            headers: getRequestHeaders()
        });
        const titles: Title[] = response.data.data.titles.map((title: any) => {
            return {
                id: title.id,
                title: title.title,
                releaseDate: title.release_date
            }
        });
        return titles;
    } catch (error) {
        console.log('error');
        throw error;
    }
}