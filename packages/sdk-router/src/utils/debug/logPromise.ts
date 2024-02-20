export async function logPromise(promise: Promise<any>, msg?: string ) {
    let result
    try {
        result = await promise

        console.log(`${msg} Promise resolved with:`)
        console.log(result)

    } catch (e) {
        console.error(`${msg} Promise rejected with: ${e}`)
    }
    return result
}

