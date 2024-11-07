import { withTokenRequest, getRequestHeaders } from '../http';
import { TeraType } from '../entity/TeraType';

export const getTeraTypes = async (): Promise<TeraType[]> => {
    try {
        const response = await withTokenRequest.get('/tera-types', {
            headers: getRequestHeaders()
        });
        const teraTypes: TeraType[] = response.data.data.tera_types.map((teraType: any) => {
            return {
                id: teraType.id,
                name: teraType.name,
            }
        });
        return teraTypes;
    } catch (error) {
        console.log('error');
        throw error;
    }
}