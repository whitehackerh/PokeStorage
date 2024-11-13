import { withTokenRequest, getRequestHeaders } from '../http';
import { Gender } from '../entity/Gender';

export const getGenders = async (): Promise<Gender[]> => {
    try {
        const response = await withTokenRequest.get('/genders', {
            headers: getRequestHeaders()
        });
        const genders: Gender[] = response.data.data.genders.map((gender: any) => {
            return {
                id: gender.id,
                name: gender.name,
            }
        });
        return genders;
    } catch (error) {
        console.log('error');
        throw error;
    }
}