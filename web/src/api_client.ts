import axios, { AxiosResponse } from 'axios';
import camelCaseKeys from "camelcase-keys";

import { EventModel } from './models/event';

export class ApiClient {
  private axiosInstance;

  constructor(private endpoint: string) {
    this.axiosInstance = axios.create({
      headers: { "Content-Type": "application/json" },
      responseType: "json",
    });
    this.axiosInstance.interceptors.response.use((res: AxiosResponse): AxiosResponse => {
      res.data = camelCaseKeys(res.data, { deep: true });
      return res;
    });
  }

  private buildUrl(path: string): string {
    return `${this.endpoint}${path}`;
  }

  async getEvents(page: number): Promise<EventModel[]> {
    const url = this.buildUrl(`/events?page=${page}&count=20`);
    const res = await this.axiosInstance.get<EventModel[]>(url);
    return res.data;
  }
}
