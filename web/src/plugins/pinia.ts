import { ServiceFactory } from "@/services"

export const piniaService = () => ({
  catalogueService: ServiceFactory.createCatalogueService('http://localhost:8080')
})
