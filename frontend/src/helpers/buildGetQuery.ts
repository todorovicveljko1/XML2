
export function omitEmpty(data: any): any{
    return Object.fromEntries(Object.entries(data).filter((v)=>!!v[1]))
 }
 
 export function buildGetQuery(data: any): string{
     return new URLSearchParams(omitEmpty(data)).toString()
 }