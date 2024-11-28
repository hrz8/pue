import type { IHttpClient } from './httpClient';

export type Room = {
  id: string
  roomType: string
  hotel: Hotel
  brand: Brand
  city: string
  prices: Price[]
  maxAdults: number
  maxChildren: number
  images: string[]
  area: number
  description: string
}

export type Hotel = {
  slug: string
  name: string
  brand: Brand
  city: string
  geo: Geo
  address: string
  website: string
  description: string
}

export type Brand = {
  slug: string
  name: string
  description: string
}

export type Geo = {
  lat: number
  lon: number
}

export type Price = {
  price: number
  symbol: string
  currency: string
}

export class CatalogueService {
  constructor(private readonly httpClient: IHttpClient) {}

  async listRooms(hotelId: string, lang = 'en-US'): Promise<{rooms: Room[]}> {
    if (hotelId === '') {
      return {rooms: []}
    }
    const response = await this.httpClient.get<{rooms: Room[]}>(`/api/hotels/${hotelId}/rooms?lang=${lang}`);
    return response.data;
  }
}
