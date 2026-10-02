<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Hash;

class AuthController extends Controller
{
    public function register(Request $request)
    {
        $data = $request->validate([
            'name' => 'required|string|max:255',
            'email' => 'required|email|unique:users,email',
            'password' => 'required|string|min:8|confirmed',
        ]);

        $user = User::create([
            'name' => $data['name'],
            'email' => $data['email'],
            'password' => Hash::make($data['password']),
            'peran' => 'mahasiswa',
        ]);

        $token = $user->createToken(
            'token-mahasiswa',
            ['mahasiswa:baca']
        )->plainTextToken;

        return response()->json([
            'message' => 'Registrasi berhasil',
            'data' => [
                'user' => $user,
                'token' => $token,
            ],
        ], 201);
    }

    public function login(Request $request)
    {
        $data = $request->validate([
            'email' => 'required|email',
            'password' => 'required|string',
        ]);

        $user = User::where('email', $data['email'])->first();

        if (!$user || !Hash::check($data['password'], $user->password)) {
            return response()->json([
                'message' => 'Email atau password salah',
            ], 401);
        }

        // Mencatat waktu login terakhir
        $user->update([
            'terakhir_login' => now(),
        ]);

        $abilities = $user->peran === 'admin'
            ? ['mahasiswa:baca', 'mahasiswa:tulis']
            : ['mahasiswa:baca'];

        $token = $user->createToken(
            'token-' . $user->peran,
            $abilities
        )->plainTextToken;

        return response()->json([
            'message' => 'Login berhasil',
            'data' => [
                'user' => $user,
                'token' => $token,
                'abilities' => $abilities,
            ],
        ]);
    }

    public function profile(Request $request)
    {
        return response()->json([
            'data' => [
                'user' => $request->user(),
                'abilities' => $request->user()->currentAccessToken()->abilities,
            ],
        ]);
    }

    public function logout(Request $request)
    {
        $request->user()->currentAccessToken()->delete();

        return response()->json([
            'message' => 'Logout berhasil',
        ]);
    }

    public function logoutSemua(Request $request)
    {
        $request->user()->tokens()->delete();

        return response()->json([
            'message' => 'Semua token berhasil dihapus',
        ]);
    }

    public function ubahPassword(Request $request)
    {
        $data = $request->validate([
            'password_lama' => 'required|string',
            'password_baru' => 'required|string|min:8|confirmed',
        ]);

        $user = $request->user();

        if (!Hash::check($data['password_lama'], $user->password)) {
            return response()->json([
                'message' => 'Password lama salah',
            ], 400);
        }

        $user->update([
            'password' => Hash::make($data['password_baru']),
        ]);

        return response()->json([
            'message' => 'Password berhasil diubah',
        ]);
    }
}