import axios from 'axios';
import { EventModel } from './models/event';

export class ApiClient {
  constructor(private endpoint: string) {
  }

  private buildUrl(path: string): string {
    return `${this.endpoint}${path}`;
  }

  async getEvents(page: number): Promise<EventModel[]> {
    const url = this.buildUrl(`/events?page=${page}&count=20`);
    const res = await axios.get<EventModel[]>(url);
    return res.data;
  }
}
