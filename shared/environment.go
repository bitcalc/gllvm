//
// OCCAM
//
// Copyright (c) 2017, SRI International
//
//  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of SRI International nor the names of its contributors may
//   be used to endorse or promote products derived from this software without
//   specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package shared

import (
	"os"
)

const (

	//ELFSectionName is the name of our ELF section of "bitcode paths".
	ELFSectionName = ".llvm_bc"

	//DarwinSegmentName is the name of our MACH-O segment of "bitcode paths".
	DarwinSegmentName = "__WLLVM"

	//DarwinSectionName is the name of our MACH-O section of "bitcode paths".
	DarwinSectionName = "__llvm_bc"
)

//LLVMToolChainBinDir is the user configured directory holding the LLVM binary tools.
var LLVMToolChainBinDir string

//LLVMCCName is the user configured name of the clang compiler.
var LLVMCCName string

//LLVMCXXName is the user configured name of the clang++ compiler.
var LLVMCXXName string

//LLVMARName is the user configured name of the llvm-ar.
var LLVMARName string

//LLVMLINKName is the user configured name of the llvm-link.
var LLVMLINKName string

//LLVMConfigureOnly is the user configured flag indicating a single pass mode is required.
var LLVMConfigureOnly string

//LLVMBitcodeStorePath is the user configured location of the bitcode archive.
var LLVMBitcodeStorePath string

//LLVMLoggingLevel is the user configured logging level: ERROR, WARNING, INFO, DEBUG.
var LLVMLoggingLevel string

//LLVMLoggingFile is the path to the optional logfile (useful when configuring)
var LLVMLoggingFile string

const (
	envpath = "LLVM_COMPILER_PATH"
	envcc   = "LLVM_CC_NAME"
	envcxx  = "LLVM_CXX_NAME"
	envar   = "LLVM_AR_NAME"
	envlnk  = "LLVM_LINK_NAME"
	envcfg  = "WLLVM_CONFIGURE_ONLY"
	envbc   = "WLLVM_BC_STORE"
	envlvl  = "WLLVM_OUTPUT_LEVEL"
	envfile = "WLLVM_OUTPUT_FILE"
)

func init() {

	LLVMToolChainBinDir = os.Getenv(envpath)
	LLVMCCName = os.Getenv(envcc)
	LLVMCXXName = os.Getenv(envcxx)
	LLVMARName = os.Getenv(envar)
	LLVMLINKName = os.Getenv(envlnk)

	LLVMConfigureOnly = os.Getenv(envcfg)
	LLVMBitcodeStorePath = os.Getenv(envbc)

	LLVMLoggingLevel = os.Getenv(envlvl)
	LLVMLoggingFile = os.Getenv(envfile)

}

func printEnvironment() {
	vars := []string{envpath, envcc, envcxx, envar, envlnk, envcfg, envbc, envlvl, envfile}

	LogWrite("\nLiving in this environment:\n\n")
	for _, v := range vars {
		LogWrite("%v = %v\n", v, os.Getenv(v))
	}

}
