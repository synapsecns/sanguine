/**
 * Type utility for improving type display in IDEs
 *
 * The Prettify type helper forces TypeScript to display a more readable version of complex types.
 * This is only done at the top level, not recursively.
 *
 * This helps with:
 * - Partial<T> - shows all properties as optional
 * - Required<T> - shows all properties as required
 * - Intersection types (A & B) - shows combined properties
 *
 * @example
 * ```ts
 * // Works with object types, intersections, and mapped types
 * type User = { name: string, age: number };
 * type PartialUser = Prettify<Partial<User>>; // Shows { name?: string, age?: number }
 *
 * // Works with intersection types
 * type A = { a: string };
 * type B = { b: number };
 * type AB = Prettify<A & B>; // Shows { a: string, b: number }
 * ```
 */
export type Prettify<T> = {
  [K in keyof T]: T[K]
} & {}
