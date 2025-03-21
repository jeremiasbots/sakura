declare module "sakura:ev3" {
    function MoveLargeMotor(port: number, speed: number, seconds: number): void;
}

declare module "sakura:http" {
    function fetch(url: string): string;
}

interface ConsoleObject {
    log(...args: unknown[]): void;
    assert(...args: unknown[]): void;
}

declare const console: ConsoleObject;