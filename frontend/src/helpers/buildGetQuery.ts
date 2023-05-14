
export function omitEmpty(data: any): any{
    return Object.fromEntries(Object.entries(data).filter((v)=>!!v[1] && (Array.isArray(v[1]) ? v[1].length != 0 : true) ))
 }
 
 export function buildGetQuery(data: any): string{
     return new URLSearchParams(omitEmpty(data)).toString()
 }