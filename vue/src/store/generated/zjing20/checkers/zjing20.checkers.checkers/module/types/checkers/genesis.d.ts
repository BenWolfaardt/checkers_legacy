import { Params } from "../checkers/params";
import { NextGame } from "../checkers/next_game";
import { StoredGame } from "../checkers/stored_game";
import { PlayerInfo } from "../checkers/player_info";
import { Leaderboard } from "../checkers/leaderboard";
import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "zjing20.checkers.checkers";
/** GenesisState defines the checkers module's genesis state. */
export interface GenesisState {
    params: Params | undefined;
    nextGame: NextGame | undefined;
    storedGameList: StoredGame[];
    playerInfoList: PlayerInfo[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    leaderboard: Leaderboard | undefined;
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
