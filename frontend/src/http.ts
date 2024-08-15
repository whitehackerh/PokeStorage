import axios from 'axios';
const baseUri = "http://127.0.0.1:8080/api";
const requestHeaders = {
    "Content-Type": "application/json",
    "Authorization": '',
};
export const multipartFormData = {
    "Content-Type": "multipart/form-data",
    "Authorization": '',
}
export const noTokenRequest = axios.create({
    baseURL:baseUri,
    headers:{"Content-Type":"application/json"}
});
export const withTokenRequest = axios.create({
    baseURL:baseUri,
})