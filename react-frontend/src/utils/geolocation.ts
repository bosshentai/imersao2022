export function getCurrentPosition(
  option?: PositionOptions,
): Promise<{ lat: number; lng: number }> {
  return new Promise((resolve, reject) => {
    navigator.geolocation.getCurrentPosition(
      (position) =>{
        resolve({
          
        })
      }
    )
  })
}
