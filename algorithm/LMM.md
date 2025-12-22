### RAG

检索增强生成（RAG） **是一种将文本检索模块与文本生成模块结合的框架，旨在提升知识密集型任务中生成回复的质量。**

![image-20251218132610266](./assets/image-20251218132610266.png)

**RAG 与纯 LLM 最大的区别，是它在推理阶段引入了外部、可更新的非参数化知识。**

传统 LLM 完全依赖训练期间固化在模型参数中的“内部记忆”，当遇到知识盲区时极易产生幻觉；而 RAG 将生成过程锚定在检索到的文档上，使模型能够基于真实证据给出回答。

更关键的是，RAG 的知识库可以随时更新——只需替换或新增文档，无需重新训练模型。这使得 RAG 能突破 LLM“知识冻结”的局限，成为一种灵活、可扩展且成本友好的知识增强手段。

#### 核心模块

典型 RAG 包含四个模块：分块（Chunking）、嵌入（Embedding）、重排序（Re-ranking）、生成（Generation）。架构如下图所示。

![img](./assets/v2-0dda8102f8477679478448e9a6df3cda_1440w.jpg)

1. **分块：**对知识源进行预处理，将长文档切分为独立的短文本块（如段落或文段）以用于索引。
2. **嵌入：**将每个文本块转化为高维向量表征，以编码其语义信息。
3. **重排序：**对初始检索的候选文本块进行二次排序，进一步提升检索精度。
4. **生成：**重排序后，选取前个（如3-5个）文本块作为生成器的最终上下文；LLM基于检索到的外部文本块生成答案或回复，将检索文本与用户查询拼接后输入模型。

#### 执行流程

![img](./assets/v2-640a361083d3d14a113eedae228e753b_1440w.jpg)

![image-20251218132758235](./assets/image-20251218132758235.png)

#### 评估

RAG 的评估是多维度的，其性能不仅依赖生成模型，也高度取决于检索流水线的质量。一个可靠的评估体系往往需要同时覆盖五大维度—**—检索准确率、生成质量、事实性、效率延迟、可扩展性**。

**检索准确率**：检索是 RAG 的信息入口，因此其质量几乎决定了最终输出的上限。如果检索不到正确文档，生成器再强也无法“无中生有”。这一维度通常关注系统是否成功找到了与用户问题真正相关的文本。

常用指标包括：

- [Recall@k](https://zhida.zhihu.com/search?content_id=267336093&content_type=Article&match_order=1&q=Recall@k&zd_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6aGlkYV9zZXJ2ZXIiLCJleHAiOjE3NjYwNDcwMTUsInEiOiJSZWNhbGxAayIsInpoaWRhX3NvdXJjZSI6ImVudGl0eSIsImNvbnRlbnRfaWQiOjI2NzMzNjA5MywiY29udGVudF90eXBlIjoiQXJ0aWNsZSIsIm1hdGNoX29yZGVyIjoxLCJ6ZF90b2tlbiI6bnVsbH0.0fkduoG3eoJtr9g_mhASILcNafZOCoa9pXKgOclKC2Y&zhida_source=entity)：查询的相关文档出现在前个检索结果中的比例；
- 平均倒数排名（[MRR](https://zhida.zhihu.com/search?content_id=267336093&content_type=Article&match_order=1&q=MRR&zd_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6aGlkYV9zZXJ2ZXIiLCJleHAiOjE3NjYwNDcwMTUsInEiOiJNUlIiLCJ6aGlkYV9zb3VyY2UiOiJlbnRpdHkiLCJjb250ZW50X2lkIjoyNjczMzYwOTMsImNvbnRlbnRfdHlwZSI6IkFydGljbGUiLCJtYXRjaF9vcmRlciI6MSwiemRfdG9rZW4iOm51bGx9.0pmiTlJrSgMCM_6IYa7wUxd6XlAodUG65J_pRRDFjcM&zhida_source=entity)）：首个相关文档排名的倒数平均值，奖励“高相关文档排名靠前”的检索结果；
- 平均准确率（[MAP](https://zhida.zhihu.com/search?content_id=267336093&content_type=Article&match_order=1&q=MAP&zd_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6aGlkYV9zZXJ2ZXIiLCJleHAiOjE3NjYwNDcwMTUsInEiOiJNQVAiLCJ6aGlkYV9zb3VyY2UiOiJlbnRpdHkiLCJjb250ZW50X2lkIjoyNjczMzYwOTMsImNvbnRlbnRfdHlwZSI6IkFydGljbGUiLCJtYXRjaF9vcmRlciI6MSwiemRfdG9rZW4iOm51bGx9.cxfR1MWBPUCm5_HrZedGH08uLnjDlfl3Xq58JPDt4jY&zhida_source=entity)）：评估所有相关文档的排序质量。

**生成质量**：生成模块负责最终回答，其质量不仅关乎语言表达是否自然，更重要的是——内容是否正确、可靠、基于证据。

评估指标主要包括：

- [精确匹配](https://zhida.zhihu.com/search?content_id=267336093&content_type=Article&match_order=1&q=精确匹配&zd_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6aGlkYV9zZXJ2ZXIiLCJleHAiOjE3NjYwNDcwMTUsInEiOiLnsr7noa7ljLnphY0iLCJ6aGlkYV9zb3VyY2UiOiJlbnRpdHkiLCJjb250ZW50X2lkIjoyNjczMzYwOTMsImNvbnRlbnRfdHlwZSI6IkFydGljbGUiLCJtYXRjaF9vcmRlciI6MSwiemRfdG9rZW4iOm51bGx9.MHrHZl2SOFZblEg_16l9krxZ3dnmnwc-pCB4yF9yk9g&zhida_source=entity)（EM）与[F1得分](https://zhida.zhihu.com/search?content_id=267336093&content_type=Article&match_order=1&q=F1得分&zd_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6aGlkYV9zZXJ2ZXIiLCJleHAiOjE3NjYwNDcwMTUsInEiOiJGMeW-l-WIhiIsInpoaWRhX3NvdXJjZSI6ImVudGl0eSIsImNvbnRlbnRfaWQiOjI2NzMzNjA5MywiY29udGVudF90eXBlIjoiQXJ0aWNsZSIsIm1hdGNoX29yZGVyIjoxLCJ6ZF90b2tlbiI6bnVsbH0.KHpX-kz-fJSG1mIkW6XR7Vv92aeocXYoGZrAJbcjLPU&zhida_source=entity)：问答任务中衡量输出与标准答案的重合度；
- [BLEU](https://zhida.zhihu.com/search?content_id=267336093&content_type=Article&match_order=1&q=BLEU&zd_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6aGlkYV9zZXJ2ZXIiLCJleHAiOjE3NjYwNDcwMTUsInEiOiJCTEVVIiwiemhpZGFfc291cmNlIjoiZW50aXR5IiwiY29udGVudF9pZCI6MjY3MzM2MDkzLCJjb250ZW50X3R5cGUiOiJBcnRpY2xlIiwibWF0Y2hfb3JkZXIiOjEsInpkX3Rva2VuIjpudWxsfQ.MQfXbIrhfHmo4lWE3LQAzVogHgkPLd21JcdHu36oEZE&zhida_source=entity)、[ROUGE](https://zhida.zhihu.com/search?content_id=267336093&content_type=Article&match_order=1&q=ROUGE&zd_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6aGlkYV9zZXJ2ZXIiLCJleHAiOjE3NjYwNDcwMTUsInEiOiJST1VHRSIsInpoaWRhX3NvdXJjZSI6ImVudGl0eSIsImNvbnRlbnRfaWQiOjI2NzMzNjA5MywiY29udGVudF90eXBlIjoiQXJ0aWNsZSIsIm1hdGNoX29yZGVyIjoxLCJ6ZF90b2tlbiI6bnVsbH0.WtSANYlztwFbbMtoHGR291ZEgnuZzwfRzWMuTVdap84&zhida_source=entity)：问答任务中衡量输出与标准答案的重合度；；
- 事实忠实性/幻觉率：判断模型输出是否真正基于检索文档，而非凭空编造。

**效率与延迟**:RAG 常用于实时问答场景，因此系统响应速度至关重要。此维度主要关注：检索耗时、生成延迟、端到端响应时间、内存 / 显存 / 计算资源占用。

**可扩展性**:随着知识库规模不断增加，系统是否还能保持稳定、准确的表现？可扩展性评估关注：索引规模扩大后是否会显著降低准确率、新文档加入时是否需要重新训练模型、系统在大规模内容下的吞吐能力、“索引规模—准确率”之间的变化规律。

**基准数据集**：为了标准化 RAG 的评估体系，研究社区构建了多个覆盖不同任务类型的评测数据集，这些基准数据集构成了当前 RAG 系统性能对比的“行业标准”，包括：

1. 开放域问答：自然问题（NQ）、TriviaQA、WebQuestions；
2. 事实核查：FEVER、AveriTeC（侧重可验证性）；
3. 多任务基准：KILT，整合了维基百科上的问答、对话、槽位填充、实体链接等任务；
4. 零样本检索：BEIR，覆盖生物医学、金融等31个零样本任务；
5. 多轮推理：MTRAG，面向需序列检索与推理的多轮对话；
6. 统一评估赛道：2024年起的TREC RAG赛道，基于MS MARCO数据集实现“检索、生成、来源支撑质量”的统一评估，核心指标为片段召回率与引用覆盖率。

![img](./assets/v2-546b62456adc0790cc510b7563f7334b_1440w.jpg)





### RoPE

**旋转位置编码（Rotary Position Embedding, RoPE）** 是一种为 Transformer 模型设计的**位置编码方案**，核心目标是让模型捕捉序列中 token 的位置信息，同时满足**绝对位置编码**和**相对位置编码**的优势，且具有良好的**长度外推性**。

它最早由论文 *《RoFormer: Enhanced Transformer with Rotary Position Embedding》* 提出，目前被广泛应用于 LLaMA、ChatGLM 等大语言模型中。

#### 核心思想

RoPE 的核心是**通过旋转操作将位置信息融入 token 的向量表示**，而非像传统的正弦余弦位置编码那样直接拼接或相加位置向量。具体来说，它利用**复数空间的旋转**来建模位置与 token 向量的关联，对于 token 的隐向量，会根据其在序列中的位置 *i*，对隐向量的成对维度进行旋转变换。

和相对位置编码相比，RoPE 具有更好的**外推性**，目前是大模型相对位置编码中应用最广的方式之一。**外推性是指大模型在训练时和预测时的输入长度不一致，导致模型的泛化能力下降的问题**。例如，如果一个模型在训练时只使用了512个 token 的文本，那么在预测时如果输入超过512个 token，模型可能无法正确处理。这就限制了大模型在处理长文本或多轮对话等任务时的效果。

#### 数学原理

**将位置信息转化为向量的旋转操作**：

- 每个位置对应一个**旋转角度**
- 通过复数域的旋转矩阵实现
- 保持相对位置关系的可解析性

1. **复数映射**

   ![image-20251210121659721](./assets/image-20251210121659721.png)

2. **旋转操作**

   ![image-20251210121720637](./assets/image-20251210121720637.png)

3. **相对位置感知**

   ![image-20251210121733617](./assets/image-20251210121733617.png)

![image-20251210122027978](./assets/image-20251210122027978.png)

| 对比维度         | 旋转位置编码（RoPE）                                   | 正弦余弦位置编码（Sinusoidal PE）                            | 可学习位置编码（Learned PE）                                 |
| ---------------- | ------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **核心原理**     | 基于复数空间旋转，将位置信息融入 token 向量的成对维度  | 利用正弦、余弦函数生成固定的位置向量，与 token 向量相加 / 拼接 | 初始化随机位置参数矩阵，通过训练学习最优位置表示             |
| **位置感知方式** | 天然捕捉**相对位置**，注意力分数直接体现 token 间距    | 仅捕捉**绝对位置**，无显式相对位置信息                       | 仅捕捉**绝对位置**，依赖模型隐式学习相对关系                 |
| **长度外推性**   | 优秀，可直接应用于训练时未见过的长序列，仅依赖相对位置 | 较差，长序列超出预设长度时，位置向量无规律可循               | 极差，无法处理超过训练最大长度的序列                         |
| **参数量**       | 无额外参数，频率参数为预设值                           | 无额外参数，位置向量由公式生成                               | 有额外参数，参数规模为 \(L_{max} \times d\)（\(L_{max}\)为最大序列长度，d为向量维度） |
| **计算复杂度**   | 中等，需对向量进行分组旋转、复数运算及还原             | 低，直接通过三角函数计算位置向量                             | 低，仅需参数矩阵查询与向量拼接 / 相加                        |
| **适用场景**     | 大语言模型（LLM）、长文本处理、需要超长上下文的任务    | 早期 Transformer 模型、机器翻译等中短序列任务                | 小模型、固定序列长度的任务（如固定长度的图像 - 文本对齐）    |
| **代表模型**     | LLaMA、ChatGLM、RoFormer                               | 原始 Transformer、BERT（早期版本）                           | BERT（主流版本）、GPT-1/2                                    |

### MoE

[混合专家模型](https://zhida.zhihu.com/search?content_id=255251579&content_type=Article&match_order=1&q=混合专家模型&zd_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6aGlkYV9zZXJ2ZXIiLCJleHAiOjE3NjU1MjE3OTUsInEiOiLmt7flkIjkuJPlrrbmqKHlnosiLCJ6aGlkYV9zb3VyY2UiOiJlbnRpdHkiLCJjb250ZW50X2lkIjoyNTUyNTE1NzksImNvbnRlbnRfdHlwZSI6IkFydGljbGUiLCJtYXRjaF9vcmRlciI6MSwiemRfdG9rZW4iOm51bGx9.KSQ0XSxGJgd2n_AfMEU3nCIHgpcAz24rR8tmBSA8rHA&zhida_source=entity) (Mixed Expert Models，简称 MoEs

![img](./assets/v2-1cdf81bb0e0862446571a0ab43a07cef_1440w.jpg)

**混合专家模型 (MoE) 的一个显著优势是它们能够在远少于稠密模型所需的计算资源下进行有效的预训练**。

作为一种基于 Transformer 架构的模型，混合专家模型主要由两个关键部分组成。

- **稀疏 [MoE 层](https://zhida.zhihu.com/search?content_id=255251579&content_type=Article&match_order=1&q=MoE+层&zd_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ6aGlkYV9zZXJ2ZXIiLCJleHAiOjE3NjU1MjE3OTUsInEiOiJNb0Ug5bGCIiwiemhpZGFfc291cmNlIjoiZW50aXR5IiwiY29udGVudF9pZCI6MjU1MjUxNTc5LCJjb250ZW50X3R5cGUiOiJBcnRpY2xlIiwibWF0Y2hfb3JkZXIiOjEsInpkX3Rva2VuIjpudWxsfQ.wFbUHlfs0vq9kNYiF4F4ShqQvXqw5-B99CEeFs8BjHM&zhida_source=entity)**: 这些层代替了传统 Transformer 模型中的前馈网络 (FFN) 层。MoE 层包含若干“专家”(例如 8 个)，每个专家本身是一个独立的神经网络。在实际应用中，这些专家通常是前馈网络 (FFN)，但它们也可以是更复杂的网络结构，甚至可以是 MoE 层本身，从而形成层级式的 MoE 结构。
- **门控网络或路由**: 这个部分用于决定哪些Token (token) 被发送到哪个专家。例如，在下图中，“More”这个Token可能被发送到第二个专家，而“Parameters”这个Token被发送到第一个专家。有时，一个Token甚至可以被发送到多个专家。Token的路由方式是 MoE 使用中的一个关键点，因为路由器由学习的参数组成，并且与网络的其他部分一同进行预训练。

![img](./assets/v2-5bd5606dd339a2318e3661925cdeb132_1440w.jpg)

尽管混合专家模型 (MoE) 提供了若干显著优势，例如更高效的预训练和与稠密模型相比更快的推理速度，但它们也伴随着一些挑战:

- **训练挑战**: 虽然 MoE 能够实现更高效的计算预训练，但它们在微调阶段往往面临泛化能力不足的问题，长期以来易于引发过拟合现象。
- **推理挑战**: MoE 模型虽然可能拥有大量参数，但在推理过程中只使用其中的一部分，这使得它们的推理速度快于具有相同数量参数的稠密模型。然而，这种模型需要将所有参数加载到内存中，因此对内存的需求非常高。以 Mixtral 8x7B 这样的 MoE 为例，需要足够的 VRAM 来容纳一个 47B 参数的稠密模型。**之所以是 47B 而不是 8 x 7B = 56B，是因为在 MoE 模型中，只有 FFN 层被视为独立的专家，而模型的其他参数是共享的**。此外，假设每个Token只使用两个专家，那么推理速度 (以 FLOPs 计算) 类似于使用 12B 模型 (而不是 14B 模型)，因为虽然它进行了 2x7B 的矩阵乘法计算，但某些层是共享的。

### LoRA

LORA(Low-Rank Adaptation，**低秩适应**)是一种模型的轻量微调方法，通过向模型的部分层添加可训练的低秩矩阵模块，实现模型在特定任务上的能力调整，同时保持原模型参数不变。由微软研究员 Edward J.Hu 等人在2021年提出，原论文为《LoRA: Low-Rank Adaptation of Large LanguageModels》

![图片](./assets/43c971fd23b39f078477aa61fb6aa89f.png)

>  核心思想：在不修改模型参数的情况下，通过添加一个额外的可训练模块，对模型行为进行小范围、有针对性的调整。
>
> - **冻结**预训练大模型的所有原始参数
> - 只在原始权重旁添加**低秩适配器**进行微调
> - 训练时只更新适配器的少量参数
>
> **微调阶段**: 仅对低秩矩阵进行参数调整在大模型中，像注意力层的权重矩阵通常规模很大。如果直接对这些矩阵进行微调，不仅计算开销大，还容易破坏原有能力。
>
> ![图片](./assets/0f7cb5b5e3fe69d745ee24bb3ff12da4.png)
>
> **推理阶段**: 在计算时叠加 LORA 模块
>
> ![图片](./assets/33895565774660e4aab34f060bac419d.png)

参数量较全参数微调（Fine-Tuning）显著降低，参数量和现有高效参数微调方法持平或更低。

基于大模型的内在低秩特性，增加旁路矩阵来模拟全参数微调，LoRA 通过简单有效的方案来达成轻量微调的目的。

为什么用LoRA？

> - 计算资源需求低
>
> 只需训练极少量参数，显著降低显存和计算开销。训练速度更快，甚至在消费级显卡上也能完成微调。
>
> - 性能接近全量参数微调
>
> 通过插入的低秩矩阵，LORA 能有效捕捉任务相关特征。在大多数 NLP 和生成任务中，其性能与全量微调相近。
>
> - 模块小巧，部署灵活，即插即用
>
> LORA 微调产生的模块通常仅几十到几百 MB，便于存储、分享与快速切换。它与原始大模型解耦，推理时加载目标 LORA 模块即可，无需改动底座模型，真正实现“即插即用”。

典型场景

>大语言模型(LLM): 通过 LORA 微调适配具体任务(如旅行助手、代码生成)，无需改动原始基座模型。
>
>扩散模型: 利用 LORA 定制化生成不同风格的图片(如艺术、写实、赛博朋克风格)。
>
>边端设备: 只需部署一个基座模型，结合多个 LORA 动态加载切换，即可适配不同推理任务，降低资源开销。
>
>![图片](./assets/e86b6a5021b386ed3d6ec73607558b13.png)





### DeepSeek

#### MLA

（多头潜在注意力 Multi-Head Latent Attention, MLA）

在讨论多头潜在注意力 (MLA) 之前，我们从[分组查询](https://so.csdn.net/so/search?q=分组查询&spm=1001.2101.3001.7020)注意力 (GQA) 说起，近年来，GQA 已成为多头注意力 (MHA) 的一种新型标准替代方案，因为它在计算和参数效率上更胜一筹。

GQA 的核心思想可以简要概括如下：与 [MHA](https://so.csdn.net/so/search?q=MHA&spm=1001.2101.3001.7020) 中每个注意力头都有自己独立的键（key）和值（value）不同，为了减少内存占用，GQA 将多个查询头分组，使其共享同一组键和值投影。例如，如图 2 所示，如果存在 2 个键值组和 4 个注意力头，那么注意力头 1 和 2 可能共享一组键值，而注意力头 3 和 4 共享另一组。这减少了键值计算的总量，从而降低了内存使用并提高了效率（根据消融研究，对模型性能没有显著影响）。

![img](./assets/cea7ebf8f286e65b03c24c28ff55ad16.png)

GQA 的核心思想是通过让多个查询头共享键（key）和值（value）来减少键值头的数量。这不仅（1）降低了模型的参数数量，而且（2）减少了推理时键值张量（tensor）的内存带宽使用，因为需要存储和从 KV 缓存中检索的键值更少。

现在，多头潜在注意力 (MLA) 提出了一种不同的内存节省策略，该策略特别适用于 KV 缓存。与 GQA 共享键（key）和值（value）头不同，MLA 在将键和值张量存储到 KV 缓存之前，将其压缩到较低维度的空间中。在推理时，这些压缩后的张量会先被恢复到原始大小，然后才能被使用，如图 3 所示。这虽然增加了额外的矩阵乘法运算，但显著减少了内存占用。

![img](./assets/acad30ae03ed3bfa11b651b4fe186620.png)

**MLA 是一种巧妙的技巧，既能减少 KV 缓存内存使用，又能在模型性能方面略微优于 MHA。**

#### MoE

专家混合（Mixture-of-Experts, MoE）

MoE 的核心思想是用**多个专家层替换 Transformer 块中的每个前馈神经网络（FeedForward）模块**，其中每个**专家层也都是一个前馈神经网络模块**。这意味着我们将单个前馈层替换为多个前馈层，如图5 所示。

![img](./assets/cb1bd5f98a2a171747a96793f7573d79.png)

Transformer 块内部的前馈神经网络（图中深灰色模块）通常包含模型总参数的很大一部分。（请注意，Transformer 块以及前馈神经网络在 LLM 中会重复很多次；在 DeepSeek-V3 中，重复了 61 次。）

因此，将**单个**前馈块替换为**多个**前馈块（如 MoE 设置中所做）会大幅增加模型的总参数数量。然而，其中的关键技巧在于，我们不会为每个词元都使用（“激活”）所有专家。相反，一个*路由*机制会为每个词元选择一小部分专家。由于一次只激活少数专家，MoE 模块通常被称为*稀疏*模块，与始终使用全部参数集的*密集*模块形成对比。然而，MoE 通过大量参数增加了 LLM 的容量，这意味着它在训练期间可以吸收更多的知识。稀疏性则保持了推理效率，因为我们不会同时使用所有参数。

例如，DeepSeek-V3 的每个 MoE 模块拥有 256 位专家，模型总参数量达到 6710 亿。但在推理时，每次只激活 9 位专家（1 位共享专家外加 8 位由路由选择的专家）。这意味着每步推理只使用 370 亿参数，而不是全部 6710 亿。DeepSeek-V3 的 MoE 设计中一个显著特点是使用了“共享专家”。这是一个始终为每个词元激活的专家。这个想法并不新鲜，它在 DeepSeek 2024 MoE 和 2022 年 DeepSpeedMoE 论文 中就已经被引入了。
![img](./assets/2560b3d34ad5020b10aa76596522e1d0.png)

在 DeepSpeedMoE 论文 中，首次指出共享专家的好处，他们发现其相比没有共享专家的情况能提升整体模型性能。这可能是因为常见的或重复的模式无需由多个独立专家分别学习，从而让这些专家有更多空间学习更特化的模式。

#### 总结
总而言之，DeepSeek-V3 是一个拥有 6710 亿参数的大型模型，发布时其性能超越了包括 4050 亿参数的 Llama3 在内的其他开源模型。尽管体量更大，但得益于其专家混合（MoE）架构，它在推理时效率更高，因为每个词元仅激活一小部分（仅 370 亿）参数。

另一个关键的显著特征是 DeepSeek-V3 采用了多头潜在注意力 (MLA) 而非分组查询注意力 (GQA)。MLA 和 GQA 都是标准多头注意力（MHA）的推理高效替代方案，尤其是在使用 KV 缓存时。尽管 MLA 的实现更为复杂，但 DeepSeek-V2 论文中的研究表明，它比 GQA 具有更好的模型性能。











