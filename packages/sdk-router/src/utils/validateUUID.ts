/**
 * Validates an UUIDv7 string (time-ordered UUID encoding Unix timestamp)
 *
 * @param uuid the generated UUIDv7 string
 */
export const validateUUID = (uuid: string): boolean => {
  const uuidv7Regex =
    /^[0-9a-f]{8}-[0-9a-f]{4}-7[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i
  return uuidv7Regex.test(uuid)
}
