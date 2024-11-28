import { CatalogueService } from './catalogue'
import { AxiosHttpClient } from '../libs/axios'

export class ServiceFactory {
  static createCatalogueService(baseURL: string) {
    const httpClient = new AxiosHttpClient(baseURL);
    return new CatalogueService(httpClient);
  }
}
