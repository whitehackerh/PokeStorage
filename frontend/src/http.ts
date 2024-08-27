import axios from 'axios';
const noTokenBaseUri = "http://127.0.0.1:8080/api";
const withTokenBaseUri = noTokenBaseUri + '/auth';
export const requestHeaders = {
    "Content-Type": "application/json",
    "Authorization": '',
};
export const multipartFormData = {
    "Content-Type": "multipart/form-data",
    "Authorization": '',
}
export const noTokenRequest = axios.create({
    baseURL: noTokenBaseUri,
    headers: {"Content-Type":"application/json"}
});
export const withTokenRequest = axios.create({
    baseURL: withTokenBaseUri,
})
export const getRequestHeaders = () => {
    requestHeaders.Authorization = `${localStorage.getItem('token_type')} ${localStorage.getItem('access_token')}`;
    return requestHeaders
}