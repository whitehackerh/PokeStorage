import { withTokenRequest, getRequestHeaders } from '../http';
import { Item } from '../entity/Item';

export const getItems = async (titleId: number): Promise<Item[]> => {
    try {
        const response = await withTokenRequest.get(`/items/${titleId}`, {
            headers: getRequestHeaders()
        });
        const items: Item[] = response.data.data.items.map((item: any) => {
            return {
                id: item.id,
                name: item.name,
            }
        });
        return items;
    } catch (error) {
        console.log('error');
        throw error;
    }
}