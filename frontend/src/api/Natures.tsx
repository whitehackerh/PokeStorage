import { withTokenRequest, getRequestHeaders } from '../http';
import { Nature } from '../entity/Nature';

export const getNatures = async (): Promise<Nature[]> => {
    try {
        const response = await withTokenRequest.get('/natures', {
            headers: getRequestHeaders()
        });
        const natures: Nature[] = response.data.data.natures.map((nature: any) => {
            return {
                id: nature.id,
                name: nature.name,
            }
        });
        return natures;
    } catch (error) {
        console.log('error');
        throw error;
    }
}